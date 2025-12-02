from tda_grafo.grafo import Grafo

def navegar_primer_link(origen, grafo: Grafo):
    recorrido = [origen]
    visitados = set()
    visitados.add(origen)

    pagina_actual = origen

    while len(recorrido) < 20:
        adyacentes = grafo.adyacentes(pagina_actual)

        if not adyacentes:
            break

        primer_link = adyacentes[0]
        recorrido.append(primer_link)

        if primer_link in visitados:
            break

        visitados.add(primer_link)
        pagina_actual = primer_link

    return recorrido
