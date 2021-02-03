//go:generate go-poor-generics g -package optional -outFile v2/generated_time.go -templateFile v2/template_for_generics.txt --headerFile v2/header_for_times.txt -namesToPrimitiveTypes "Time=time.Time,Duration=time.Duration"
package optional
