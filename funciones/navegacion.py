from tda_grafo.grafo import Grafo


# Navega desde una página siguiendo siempre el primer link disponible.
# Se para cuando no hay más links, se completan 20 páginas o hay un ciclo
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
        
        if primer_link in visitados:
            recorrido.append(primer_link)
            break
        
        recorrido.append(primer_link)
        visitados.add(primer_link)
        pagina_actual = primer_link
    
    return recorrido
