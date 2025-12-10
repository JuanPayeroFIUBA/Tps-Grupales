from tda_grafo.grafo import Grafo


# Funcion que ejecuta el comando "navegación", que navega desde una página siguiendo siempre el primer link disponible y se para cuando no hay más links o se completan 20 páginas
def navegar_primer_link(origen, grafo: Grafo):
    recorrido = [origen]
    pagina_actual = origen

    while len(recorrido) < 20:
        adyacentes = grafo.adyacentes(pagina_actual)

        if not adyacentes:
            break

        primer_link = adyacentes[0]

        recorrido.append(primer_link)
        pagina_actual = primer_link

    return recorrido
