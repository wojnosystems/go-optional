//go:generate go-poor-generics g -package optional -outFile generated_time.go -templateFile template_for_generics.txt --headerFile header_for_times.txt -namesToPrimitiveTypes "Time=time.Time,Duration=time.Duration"
package optional
