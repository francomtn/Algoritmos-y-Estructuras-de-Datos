package abb

type ab struct {
	izq   *ab
	der   *ab
	clave string
	dato  int
}

func ReconstruirParticular() *ab {
	PRE_ORDER := "EURMAONDVSZT"
	IN_ORDER := "MRAUOZSVDNET"
	ab := _reconstruir(PRE_ORDER, IN_ORDER)
	return ab
}

func _reconstruir(preOrder, inOrder string) *ab {

	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	ab := &ab{clave: string(preOrder[0])}

	indInOrder := buscarIndice(inOrder, ab.clave)
	ab.izq = _reconstruir(preOrder[1:indInOrder+1], inOrder[:indInOrder])
	ab.der = _reconstruir(preOrder[indInOrder+1:], inOrder[indInOrder+1:])
	return ab
}

func buscarIndice(inOrd, clave string) int {

	pos := 0
	for ind, elem := range inOrd {
		if string(elem) == clave {
			pos = ind
			break
		}
	}
	return pos
}

/*
package reconstruccion

const(
    PRE_ORDER = "EURMAONDVSZT"
    IN_ORDER  = "MRAUOZSVDNET"
)

type ab struct {
    izq *ab
    der *ab
    clave string
}

func buscarin_order(elem, in_order string) int{
    var pos_in int

    for i, v := range in_order {
        if string(v) == elem{
            pos_in = i
            break
        }
    }

    return pos_in
}

func ReconstruirParticular(pre_order, in_order string) *ab {

    if len(pre_order) == 0 || len(in_order) == 0 {
        return nil
    }

    ab := &ab{clave: string(pre_order[0])}

    pos_in := buscarin_order(ab.clave, in_order)

    ab.izq = ReconstruirParticular(pre_order[1:pos_in+1], in_order[:pos_in])
    ab.der = ReconstruirParticular(pre_order[pos_in+1:], in_order[pos_in+1:])

    return ab
}*/
