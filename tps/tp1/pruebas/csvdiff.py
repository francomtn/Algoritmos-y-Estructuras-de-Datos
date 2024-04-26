from math import isclose
from sys import argv, stdin, stderr


def eprintf(s, *args, **kwargs):
    print(s.format(*args, **kwargs), file=stderr)
    

def isfloat(value):
    try:
        float(value)
        return True
    except:
        return False


def acortar(linea, maxlen=120):
    if len(linea) > maxlen:
        linea = linea[:maxlen+1] + " […]"
    return linea


def report_diff(line_no, esperada, obtenida):
    esperada = acortar(esperada)
    obtenida = acortar(obtenida)
    eprintf("{}: diferencia en línea número {}: esperado '{}', obtenido '{}'.",
        argv[1],
        line_no,
        esperada.strip() if esperada else "(fin de archivo)",
        obtenida.strip() if obtenida else "(fin de archivo)"
    )


def line_diff(line_no, l_esperada, l_obtenida):
    """ Reporta el diff entre dos líneas de distintos archivos por
    stderr.
    
    Si alguna de las dos es empty string es porque alguno de los dos
    archivos ha terminado, y devuelve False. """
    esperados = l_esperada.strip().split(",")
    obtenidos = l_obtenida.strip().split(",")
    
    if len(esperados) != len(obtenidos):
        report_diff(line_no, str(esperados), str(obtenidos))
        exit(1)
        
    for i in range(len(esperados)):
        f1 = esperados[i]
        f2 = obtenidos[i]
        
        # En caso de que los dos campos sean flotantes se los compara.
        if isfloat(f1) and isfloat(f2):
            f1 = float(f1)
            f2 = float(f2)
            
            if not isclose(f1, f2, abs_tol=0.5):
                report_diff(line_no, l_esperada, l_obtenida)
                exit(1)
        
        # Si sólo uno es flotante, hay un error.
        elif isfloat(f1) or isfloat(f2):
            report_diff(line_no, l_esperada, l_obtenida)
            exit(1)
        
        # Si ninguno es flotante se comparan los strings.
        else:
            if f1 != f2:
                report_diff(line_no, l_esperada, l_obtenida)
                exit(1)


def csv_diff(f_esperado, f_obtenido):
    line_no = 1   # Para reportar errores.
    while True:
        l1 = f_esperado.readline()
        l2 = f_obtenido.readline()
      
        # Los dos archivos terminan en simultáneo, todo bien.
        if not l1 and not l2:
            exit(0)
            
        # Un archivo termina y el otro no, es un error.
        if not l1 or not l2:
            report_diff(line_no, l1, l2)
            exit(1)

        # Ningún archivo terminó aún, se calcula el diff.
        line_diff(line_no, l1, l2)
        
        line_no += 1


if __name__ == "__main__":
    if len(argv) != 3:
        eprintf("Esperados dos parámetros.")
        exit(-1)
    
    # Compara contra argv[1] contra stdin si argv[2] es "-" (como diff).
    if argv[2] == "-":
        with open(argv[1]) as fp1:
            csv_diff(fp1, stdin)
        
    else:
        with open(argv[1]) as fp1, open(argv[2]) as fp2:
            csv_diff(fp1, fp2)
