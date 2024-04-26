#!/usr/bin/env bash

set -eu
#
# Uso: ./pruebas.sh <RUTA-PROGRAMA> [PRUEBAS.yaml]
# Para alumnos, unicamente llamar: ./pruebas.sh <RUTA-PROGRAMA>

CSVDIFF="csvdiff.py"
YAMLTAP="../yamltap.py"

PROGRAMA="$1"
BASENAME=$(basename "$PROGRAMA")

if [[ $# -ge 2 ]]; then
    $YAMLTAP --gen-only $2 $PROGRAMA
fi

RET=0

echo "Ejecución de pruebas unitarias de $BASENAME:"
echo ""
for t in *.test; do
    b=${t%.test}
    echo "$b $(< $t)"
    ($PROGRAMA <${b}_in >tmp_out 2>tmp_err     && \
            python3 $CSVDIFF ${b}_out tmp_out && \
            python3 $CSVDIFF ${b}_err tmp_err && \
            echo "OK") || { RET=$?; echo "ERROR"; }
done

rm tmp_out tmp_err

exit $RET