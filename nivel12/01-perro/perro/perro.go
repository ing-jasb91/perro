// Package perro permite entender un poco a los perros
package perro

// Edad recibe los años humanos como valor de tipo int y retorna la edad en años perro
func Edad(e int) int {
	ap := 7 * e
	return ap
}
