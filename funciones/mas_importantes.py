from tda_grafo.grafo import Grafo
import heapq


# Funcion encargada de ejecutar el comando "mas_importantes" que nos muestra las n páginas más centrales/importantes del mundo según el algoritmo de pagerank, ordenadas de mayor importancia a menor importancia
def calcular_mas_importantes(n, grafo: Grafo):
    """
    Calcula PageRank iterativo para encontrat las n paginas mas importantes.
    se simula alguien que navega entre links con probabilidad d
    o salta a cualquier pagina con probabilidad (1-d)
    """
    vertices = grafo.obtener_vertices()
    cantidad_vertices = len(vertices)

    if cantidad_vertices == 0:
        return []

    if n > cantidad_vertices:
        n = cantidad_vertices

    # probabilidad de seguir un link 85%
    d = 0.85
    iteraciones = 20  # cantidad de iteraciones max

    # Se inicializa todos con el mismo pagerank
    pagerank = {}
    for v in vertices:
        pagerank[v] = 1 / cantidad_vertices

    # se itera hasta converger o llegar al límite
    for _ in range(iteraciones):
        nuevo_pagerank = {}

        for v in vertices:
            # se empieza con la probabilidad de salto aleatorio
            suma = (1 - d) / cantidad_vertices

            # cada pagina que apunta a v le pasa una parte de su importancia
            """ESTE TERCER FOR NO HACE QUE LA COMPLEJIDAD SEA DE O(P^2)??"""
            for u in vertices:
                if grafo.estan_unidos(u, v):
                    salientes_u = len(grafo.adyacentes(u))
                    if salientes_u > 0:
                        suma += d * pagerank[u] / salientes_u

            nuevo_pagerank[v] = suma

        pagerank = nuevo_pagerank

    # Usamos un heap para tener los n más importantes
    heap_mas_importantes = []

    for pagina, valor in pagerank.items():
        if len(heap_mas_importantes) < n:
            heapq.heappush(heap_mas_importantes, (valor, pagina))
        else:
            if valor > heap_mas_importantes[0][0]:
                heapq.heapreplace(heap_mas_importantes, (valor, pagina))

    # ordenamos de mayor a menor
    resultado = sorted(heap_mas_importantes, key=lambda x: -x[0])

    return [pagina for _, pagina in resultado]
