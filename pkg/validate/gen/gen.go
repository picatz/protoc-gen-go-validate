package gen

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/picatz/protoc-gen-go-validate/pkg/validate"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type Generator struct {
	pb *protogen.Plugin
}

func NewGenerator(pb *protogen.Plugin) *Generator {
	return &Generator{pb: pb}
}

func (g *Generator) Generate() error {
	gp := g.pb
	for _, name := range gp.Request.FileToGenerate {
		f := gp.FilesByPath[name]

		if len(f.Messages) == 0 {
			glog.V(1).Infof("Skipping %s, no messages", name)
			continue
		}

		fileName := fmt.Sprintf("%s.pb.validate.go", f.GeneratedFilenamePrefix)
		gf := gp.NewGeneratedFile(fileName, f.GoImportPath)
		gf.P(fmt.Sprintf("package %s", f.GoPackageName))

		fmtErrorf := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "fmt", GoName: "Errorf"})
		glog.V(0).Infof("using %v", fmtErrorf)

		for _, msg := range f.Messages {
			// generate validation function for message
			gf.P(`// Validate applies configured validation rule options from the protobuf.`)
			gf.P(fmt.Sprintf(`func (x *%s) Validate() error {`, msg.GoIdent.GoName))
			var anyValidations bool
			for _, field := range msg.Fields {
				name := field.Desc.Name()

				glog.V(0).Infof("field: %v", name)

				fOpts, ok := field.Desc.Options().(*descriptorpb.FieldOptions)
				if !ok {
					return fmt.Errorf("failed to cast %T to *descriptorpb.FieldOptions", field.Desc.Options())
				}

				v := proto.GetExtension(fOpts, validate.E_Field)

				opts, ok := v.(*validate.FieldRules)
				if !ok {
					return nil
				}

				glog.V(0).Infof("\toptions: %v", opts)

				if opts == nil {
					continue
				}

				if opts.Message != nil {
					if opts.Message.Skip != nil && opts.Message.GetSkip() {
						glog.V(0).Infof("skipping validation for field %q", name)
						continue
					}

					if opts.Message.Required != nil && opts.Message.GetRequired() {
						anyValidations = true
						gf.P(fmt.Sprintf(`    if x.%s == nil {`, field.GoName))
						gf.P(fmt.Sprintf(`        return %s("invalid value for %s, cannot be nil")`, fmtErrorf, field.Desc.Name()))
						gf.P(`                }`)
					}
				}

				switch opts.Type.(type) {
				case *validate.FieldRules_String_:
					// spot check human error
					if field.Desc.Kind().String() != "string" {
						return fmt.Errorf("invalid validation string rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetString_()
					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						if rules.Required != nil && rules.GetRequired() {
							gf.P(fmt.Sprintf(`    if len(%s) == 0 {`, accessName))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid value for %s, cannot be empty")`, field.GoName))
							gf.P(`                }`)
						}

						//if field.Desc.HasOptionalKeyword() {
						//	gf.P(fmt.Sprintf(`    if x.%s != nil {`, field.GoName))
						//	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						//}

						if rules.Len != nil {
							gf.P(fmt.Sprintf(`    if len(%s) != %d {`, accessName, rules.GetLen()))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid length for %s, must be %d")`, field.GoName, rules.GetLen()))
							gf.P(`                }`)
						}
						if rules.Gt != nil {
							gf.P(fmt.Sprintf(`    if len(%s) <= %d {`, accessName, rules.GetGt()))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid length for %s, must be greater than %d")`, field.GoName, rules.GetGt()))
							gf.P(`                }`)
						}
						if rules.Lt != nil {
							gf.P(fmt.Sprintf(`    if len(%s) >= %d {`, accessName, rules.GetLt()))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid length for %s, must be less than %d")`, field.GoName, rules.GetLt()))
							gf.P(`                }`)
						}
						if rules.Min != nil {
							gf.P(fmt.Sprintf(`    if len(%s) < %d {`, accessName, rules.GetMin()))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid length for %s, must be at least %d")`, field.GoName, rules.GetMin()))
							gf.P(`                }`)
						}
						if rules.Max != nil {
							gf.P(fmt.Sprintf(`    if len(%s) > %d {`, accessName, rules.GetMax()))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid length for %s, cannot be more than %d")`, field.GoName, rules.GetMax()))
							gf.P(`                }`)
						}
						if rules.Contains != nil {
							stringContains := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings", GoName: "Contains"})

							gf.P(fmt.Sprintf(`    if !%s(%s, %q) {`, stringContains, accessName, rules.GetContains()))

							msg := fmt.Sprintf("invalid value for %s, must contain %q", field.GoName, rules.GetContains())
							gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.NotContains != nil {
							stringContains := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings", GoName: "Contains"})

							gf.P(fmt.Sprintf(`    if %s(%s, %q) {`, stringContains, accessName, rules.GetNotContains()))
							msg := fmt.Sprintf("invalid value for %s, must not contain %q", field.GoName, rules.GetNotContains())
							gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.Prefix != nil {
							stringHasPrefix := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings", GoName: "HasPrefix"})
							gf.P(fmt.Sprintf(`    if !%s(%s, %q) {`, stringHasPrefix, accessName, rules.GetPrefix()))
							msg := fmt.Sprintf("invalid value for %s, must have prefix %q", field.GoName, rules.GetPrefix())
							gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.Suffix != nil {
							stringHasSuffix := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings", GoName: "HasSuffix"})
							gf.P(fmt.Sprintf(`    if !%s(%s, %q) {`, stringHasSuffix, accessName, rules.GetSuffix()))
							msg := fmt.Sprintf("invalid value for %s, must have suffix %q", field.GoName, rules.GetSuffix())
							gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.AllowSpace != nil && !rules.GetAllowSpace() {
							stringContains := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "strings", GoName: "Contains"})

							gf.P(fmt.Sprintf(`    if %s(%s, " ") {`, stringContains, accessName))
							msg := fmt.Sprintf("invalid value for %s, cannot have spaces", field.GoName)
							gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.GetAsciiOnly() {
							gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "unicode"})
							gf.P(fmt.Sprintf(`    for _, c := range %s {`, accessName))
							gf.P(`                    if c > unicode.MaxASCII {`)
							msg := fmt.Sprintf("invalid value for %s, can only contain ASCII characters", field.GoName)
							gf.P(fmt.Sprintf(`            return fmt.Errorf(%q)`, msg))
							gf.P(`                    }`)
							gf.P(`                }`)
						}
						if rules.Match != nil {
							gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "regexp"})
							gf.P(fmt.Sprintf(`    match%s, err := regexp.Match(%q, []byte(%s))`, field.GoName, rules.GetMatch(), accessName))
							gf.P(`                if err != nil {`)
							gf.P(`                    return fmt.Errorf("failed to validate: %w", err)`)
							gf.P(`                }`)
							gf.P(fmt.Sprintf(`                if !match%s {`, field.GoName))
							msg := fmt.Sprintf("invalid value for %s, did not match %q", field.GoName, rules.GetMatch())
							gf.P(fmt.Sprintf(`            return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if rules.NotMatch != nil {
							gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "regexp"})
							gf.P(fmt.Sprintf(`    notMatch%s, err := regexp.Match(%q, []byte(%s))`, field.GoName, rules.GetNotMatch(), accessName))
							gf.P(`                if err != nil {`)
							gf.P(`                    return fmt.Errorf("failed to validate: %w", err)`)
							gf.P(`                }`)
							gf.P(fmt.Sprintf(`                if notMatch%s {`, field.GoName))
							msg := fmt.Sprintf("invalid value for %s, can not match %q", field.GoName, rules.GetNotMatch())
							gf.P(fmt.Sprintf(`            return fmt.Errorf(%q)`, msg))
							gf.P(`                }`)
						}
						if field.Desc.IsList() {
							gf.P(`}`)
						}
						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Bytes:
					// spot check human error
					if field.Desc.Kind().String() != "bytes" {
						return fmt.Errorf("invalid validation bytes rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetBytes()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						if rules.Required != nil && rules.GetRequired() {
							gf.P(fmt.Sprintf(`    if x.Get%s() == nil {`, field.GoName))
							gf.P(fmt.Sprintf(`        return fmt.Errorf("invalid value for %s, is required")`, field.GoName))
							gf.P(`                }`)
						}

						if field.Desc.HasOptionalKeyword() {
							gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
							accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						}

						if rules.Len != nil {
							val := rules.GetLen()
							expr := fmt.Sprintf(`len(%s) != %d`, accessName, val)
							errMsg := fmt.Sprintf("must equal %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`len(%s) <= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`len(%s) < %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`len(%s) >= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`len(%s) > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`len(%s) > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`len(%s) < %d`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Prefix != nil {
							val := rules.GetPrefix()
							expr := fmt.Sprintf(`strings.HasPrefix(string(%s), %q)`, accessName, val)
							errMsg := fmt.Sprintf("must have prefix %q", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Suffix != nil {
							val := rules.GetSuffix()
							expr := fmt.Sprintf(`strings.HasSuffix(string(%s), %q)`, accessName, val)
							errMsg := fmt.Sprintf("must have suffix %q", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf("len(%s) == 0", accessName)
								errMsg := "must use non-empty value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						if field.Desc.HasOptionalKeyword() {
							gf.P(`                }`)
						}
					}
				case *validate.FieldRules_Uint32:
					// spot check human error
					if field.Desc.Kind().String() != "uint32" {
						return fmt.Errorf("invalid validation uint32 rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetUint32()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %d`, accessName, val)
							errMsg := fmt.Sprintf("must equal %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Int32:
					// spot check human error
					if field.Desc.Kind().String() != "int32" {
						return fmt.Errorf("invalid validation int32 rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetInt32()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %d`, accessName, val)
							errMsg := fmt.Sprintf("must equal %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Uint64:
					// spot check human error
					if field.Desc.Kind().String() != "uint64" {
						return fmt.Errorf("invalid validation uint64 rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetUint64()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %d`, accessName, val)
							errMsg := fmt.Sprintf("must equal %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Int64:
					// spot check human error
					if field.Desc.Kind().String() != "int64" {
						return fmt.Errorf("invalid validation int64 rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetInt64()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %d`, accessName, val)
							errMsg := fmt.Sprintf("must equal %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %d`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %d`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %d", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Float:
					// spot check human error
					if field.Desc.Kind().String() != "float" {
						return fmt.Errorf("invalid validation float rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetFloat()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %f`, accessName, val)
							errMsg := fmt.Sprintf("must equal %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %f`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %f`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %f`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %f`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %f`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %f`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				case *validate.FieldRules_Double:
					// spot check human error
					if field.Desc.Kind().String() != "double" {
						return fmt.Errorf("invalid validation double rule kind used for field %q of kind %q", field.Desc.Name(), field.Desc.Kind())
					}

					anyValidations = true

					rules := opts.GetDouble()

					if rules != nil {
						accessName := fmt.Sprintf("x.Get%s()", field.GoName)

						if field.Desc.IsList() {
							gf.P(fmt.Sprintf(`for _, v := range %s {`, accessName))
							accessName = "v"
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(fmt.Sprintf(`    if %s != nil {`, accessName))
						// 	accessName = fmt.Sprintf("x.Get%s()", field.GoName)
						// }

						if rules.Eq != nil {
							val := rules.GetEq()
							expr := fmt.Sprintf(`%s != %f`, accessName, val)
							errMsg := fmt.Sprintf("must equal %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gt != nil {
							val := rules.GetGt()
							expr := fmt.Sprintf(`%s <= %f`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Gte != nil {
							val := rules.GetGte()
							expr := fmt.Sprintf(`%s < %f`, accessName, val)
							errMsg := fmt.Sprintf("must be greater than or equal to %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lt != nil {
							val := rules.GetLt()
							expr := fmt.Sprintf(`%s >= %f`, accessName, val)
							errMsg := fmt.Sprintf("must be less than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Lte != nil {
							val := rules.GetLte()
							expr := fmt.Sprintf(`%s > %f`, accessName, val)
							errMsg := fmt.Sprintf("must be less than or equal to %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Min != nil {
							val := rules.GetMin()
							expr := fmt.Sprintf(`%s > %f`, accessName, val)
							errMsg := fmt.Sprintf("must be at least %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Max != nil {
							val := rules.GetMax()
							expr := fmt.Sprintf(`%s < %f`, accessName, val)
							errMsg := fmt.Sprintf("cannot be greater than %f", val)
							writeFieldValidation(gf, field.GoName, expr, errMsg)
						}

						if rules.Required != nil {
							val := rules.GetRequired()
							if val {
								expr := fmt.Sprintf(`%s == 0`, accessName)
								errMsg := "must use non-zero value"
								writeFieldValidation(gf, field.GoName, expr, errMsg)
							}
						}

						if field.Desc.IsList() {
							gf.P(`}`)
						}

						// if field.Desc.HasOptionalKeyword() {
						// 	gf.P(`                }`)
						// }
					}
				default:
					// temporary hack
					if field.Desc.Kind().String() != "message" {
						glog.V(0).Infof("unhandled validation type: %T for field %q of kind %q", opts.Type, field.Desc.Name(), field.Desc.Kind())
					}
					continue
				}
			}
			if !anyValidations {
				gf.P(`    return fmt.Errorf("no validation options configured") // has no validations`)
				gf.P(`}`)
				gf.P("")
			} else {
				gf.P(`    return nil // is valid`)
				gf.P(`}`)
				gf.P("")
			}
		}
	}
	return nil

}

func writeFieldValidation(gf *protogen.GeneratedFile, field, expr, errMsg string) {
	gf.P(fmt.Sprintf(`    if %s {`, expr))
	errMsg = fmt.Sprintf("invalid value for %s, %s", field, errMsg)
	gf.P(fmt.Sprintf(`        return fmt.Errorf(%q)`, errMsg))
	gf.P(`                }`)

}
