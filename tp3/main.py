from locale import atoi
from funciones.ciclo import obtener_ciclo
from funciones.diametro import obtener_diametro
from funciones.camino import camino_minimo_origen_destino
from funciones.lectura import orden_topologico_lectura
import tda_grafo.grafo as g

RUTA_ARCHIVO = "wiki_parseado.txt"


def main():
    grafo = crear_grafo_desde_archivo(RUTA_ARCHIVO)
    procesar_entradas(grafo)
    return


def crear_grafo_desde_archivo(ruta):
    grafo = g.Grafo(esdirigido=True)

    with open(ruta, "r", encoding="utf-8") as archivo:
        for linea in archivo:
            linea = linea.strip()

            if not linea:
                continue

            partes = linea.split("\t")
            grafo.agregar_vertice(partes[0])

            for arista in partes[1:]:
                arista = arista.strip()
                # if arista:
                grafo.agregar_arista(partes[0], arista)
    return grafo


def procesar_entradas(grafo):
    comando = ""
    parametros = ""

    # no me acuerdo como hacer para leer entrada estandar/archivos xd

    if comando == "lista_comandos":
        return

    elif comando == "camino":
        origen, destino = parametros[1].split(",")
        if origen not in grafo or destino not in grafo:
            print("No se encontro recorrido")
        camino = camino_minimo_origen_destino(origen, destino, grafo)
        if not camino:
            print("No se encontro recorrido")
        else:
            print(" -> ".join(camino))
            print(f"Costo: {len(camino) - 1}")

    elif comando == "lectura":
        vertices = set(parametros[1].split(","))
        # podriamos chequear que todas las paginas esten en el grafo
        camino = orden_topologico_lectura(vertices, grafo)
        if not camino:
            print("No existe forma de leer las paginas en orden")
        else:
            print(", ".join(camino))

    elif comando == "diametro":
        camino, diametro = obtener_diametro(grafo)
        print(" -> ".join(camino))
        print(f"Costo: {diametro}")

    elif comando == "ciclo":
        origen, largo_str = parametros[1].split(",")
        largo = int(largo_str)
        ciclo = obtener_ciclo(origen, largo, grafo)

        if not ciclo:
            print("No se encontro recorrido")
        else:
            print(" -> ".join(ciclo))
