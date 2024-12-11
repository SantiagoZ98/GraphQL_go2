package graph

import (
	"context"
)

// Resolver struct
type Resolver struct{}

// Función resolver para la consulta holaMundo
func (r *Resolver) HolaMundo(ctx context.Context) (string, error) {
	return "¡Hola Mundo desde GraphQL en Go!", nil
}

// Genera el esquema con el resolver
func NewResolver() *Resolver {
	return &Resolver{}
}
