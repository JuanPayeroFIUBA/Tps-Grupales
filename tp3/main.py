import sys
from funciones.ciclo import obtener_ciclo
from funciones.diametro import obtener_diametro
from funciones.camino import camino_minimo_origen_destino
from funciones.lectura import orden_topologico_lectura
from funciones.rango import obtener_paginas_en_rango
from funciones.navegacion import navegar_primer_link
from funciones.mas_importantes import calcular_mas_importantes
import tda_grafo.grafo as g

COMANDO_LISTAR = "listar_operaciones"
LISTA_COMANDOS = (
    "camino \nmas_importantes \nciclo \nlectura \ndiametro \nrango \nnavegación"
)

COMANDO_CAMINO = "camino"
COMANDO_MAS_IMPORTANTES = "mas_importantes"
COMANDO_CICLOS = "ciclo"
COMANDO_LECTURA = "lectura"
COMANDO_DIAMETRO = "diametro"
COMANDO_RANGO = "rango"
COMANDO_NAVEGACION = "navegación"

ERROR_CAMINO_NO_ENCONTRADO = "No se encontro recorrido"
ERROR_LECTURA = "No existe forma de leer las paginas en orden"
ERROR_COMANDO_INVALIDO = "Comando invalido"


def main():
    if len(sys.argv) < 2:
        print("Uso: ./netstats <archivo_wiki>")
        return

    ruta_archivo = sys.argv[1]
    grafo = crear_grafo_desde_archivo(ruta_archivo)
    procesar_entradas(grafo)
    return


# Funcion encargada de crear un TDA grafo a partir del archivo pasado por args,para asi poder ejecutar los comandos del programa
def crear_grafo_desde_archivo(ruta):
    grafo = g.Grafo(dirigido=True)

    with open(ruta, "r", encoding="utf-8") as archivo:
        for linea in archivo:
            linea = linea.strip()

            if not linea:
                continue

            partes = linea.split("\t")
            pagina = partes[0]
            grafo.agregar_vertice(pagina)

    with open(ruta, "r", encoding="utf-8") as archivo:
        for linea in archivo:
            linea = linea.strip()

            if not linea:
                continue

            partes = linea.split("\t")
            pagina = partes[0]

            for link in partes[1:]:
                link = link.strip()
                if link and grafo.hay_vertice(link):
                    grafo.agregar_arista(pagina, link)

    return grafo


# Funcion encargada de procesar la informacion ingresada por entrada estandar y desglosar el comando en cuestion, y los parametros de este, para luego procesar el comando y llamar a la funcionalidad correspondiente
def procesar_entradas(grafo):
    for linea in sys.stdin:
        linea = linea.strip()
        if not linea:
            continue

        partes = linea.split(" ", 1)
        comando = partes[0]
        parametros = partes[1] if len(partes) > 1 else ""

        if comando == COMANDO_LISTAR:
            print(LISTA_COMANDOS)

        elif comando == COMANDO_CAMINO:
            origen, destino = parametros.split(",")
            if not grafo.hay_vertice(origen) or not grafo.hay_vertice(destino):
                print(ERROR_CAMINO_NO_ENCONTRADO)
                continue

            camino = camino_minimo_origen_destino(origen, destino, grafo)
            if not camino:
                print(ERROR_CAMINO_NO_ENCONTRADO)
            else:
                print(" -> ".join(camino))
                print(f"Costo: {len(camino) - 1}")

        elif comando == COMANDO_LECTURA:
            paginas = parametros.split(",")
            vertices = set(paginas)

            camino = orden_topologico_lectura(vertices, grafo)
            if not camino:
                print(ERROR_LECTURA)
            else:
                print(", ".join(camino))

        elif comando == COMANDO_DIAMETRO:
            camino, diametro = obtener_diametro(grafo)
            print(" -> ".join(camino))
            print(f"Costo: {diametro}")

        elif comando == COMANDO_CICLOS:
            partes_ciclo = parametros.split(",")
            origen = partes_ciclo[0]
            largo = int(partes_ciclo[1])

            if not grafo.hay_vertice(origen):
                print(ERROR_CAMINO_NO_ENCONTRADO)
                continue

            ciclo = obtener_ciclo(origen, largo, grafo)

            if not ciclo:
                print(ERROR_CAMINO_NO_ENCONTRADO)
            else:
                print(" -> ".join(ciclo))

        elif comando == COMANDO_RANGO:
            partes_rango = parametros.split(",")
            pagina = partes_rango[0]
            n = int(partes_rango[1])

            if not grafo.hay_vertice(pagina):
                print("0")
                continue

            cantidad = obtener_paginas_en_rango(pagina, n, grafo)
            print(cantidad)

        elif comando == COMANDO_NAVEGACION:
            pagina_origen = parametros.strip()

            if not grafo.hay_vertice(pagina_origen):
                print(pagina_origen)
                continue

            recorrido = navegar_primer_link(pagina_origen, grafo)
            print(" -> ".join(recorrido))

        elif comando == COMANDO_MAS_IMPORTANTES:
            n = int(parametros.strip()) if parametros else 20

            paginas_importantes = calcular_mas_importantes(n, grafo)
            print(", ".join(paginas_importantes))
        else:
            print(ERROR_COMANDO_INVALIDO)


if __name__ == "__main__":
    main()
