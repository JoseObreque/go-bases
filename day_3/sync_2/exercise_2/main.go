package main

import (
	"errors"
	"fmt"
)

const (
	pequeno = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

func main() {
	producto, err := Factory(1000, mediano)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Producto:", producto)
		fmt.Println("Precio:", producto.Precio())
	}
}

type Producto interface {
	Precio() float64
}

type ProductoPequeno struct {
	Costo float64
}

type ProductoMediano struct {
	Costo         float64
	Mantenimiento float64
}

type ProductoGrande struct {
	Costo         float64
	Mantenimiento float64
	CostoEnvío    float64
}

// Función que calcula el precio final de un producto pequeño en base a su costo.
func (p ProductoPequeno) Precio() float64 {
	return p.Costo
}

/*
Función que calcula el precio final de un producto mediano en base a su costo y
su mantenimiento en bodega.
*/
func (p ProductoMediano) Precio() float64 {
	return p.Costo * (1 + p.Mantenimiento)
}

/*
Función que calcula el precio final de un producto grande en base a su costo,
su mantenimiento en bodega y un costo de envío fijo.
*/
func (p ProductoGrande) Precio() float64 {
	return p.Costo*(1+p.Mantenimiento) + p.CostoEnvío
}

/*
Función que genera una instancia de un producto utilizando su costo unitario y su tamaño,
que puede ser pequeño, mediano o grande.
*/
func Factory(costoProducto float64, tipoProducto string) (Producto, error) {
	switch tipoProducto {
	case pequeno:
		producto := ProductoPequeno{
			Costo: costoProducto,
		}
		return producto, nil
	case mediano:
		producto := ProductoMediano{
			Costo:         costoProducto,
			Mantenimiento: 0.03,
		}
		return producto, nil
	case grande:
		producto := ProductoGrande{
			Costo:         costoProducto,
			Mantenimiento: 0.06,
			CostoEnvío:    2500,
		}
		return producto, nil
	default:
		return nil, errors.New("Tipo de producto no válido")
	}
}
