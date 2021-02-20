from urllib.request import urlopen
from bs4 import BeautifulSoup
import json
import re

html = urlopen("http://obelisk.me.uk/6502/reference.html")
bsObj = BeautifulSoup(html, "html.parser")

tables = bsObj.findAll("table")
h3s = bsObj.findAll("h3")


def convertMode(originalMode):
    trimMode = re.sub('(\n|\r\n|\r)\s*', ' ', originalMode.strip())

    mapping = {
        "Absolute": "Absolute",
        "Absolute,X": "AbsoluteX",
        "Absolute,Y": "AbsoluteY",
        "Accumulator": "Accumulator",
        "Immediate": "Immediate",
        "Implied": "Implied",
        "(Indirect,X)": "IndexedIndirect",
        "Indirect": "Indirect",
        "(Indirect),Y": "IndirectIndexed",
        "Relative": "Relative",
        "Zero Page": "ZeroPage",
        "Zero Page,X": "ZeroPageX",
        "Zero Page,Y": "ZeroPageY",
    }

    return mapping[trimMode]


def convertOpcode(originalOpcode):
    return re.sub('\$', '0x', originalOpcode.strip())


def convertBytes(originalBytes):
    return originalBytes.strip()


def convertCycles(originalOpcode):
    return [
        originalOpcode.strip()[0],
        1 if originalOpcode.strip().find("crossed") != -1 else 0,
    ]


index = 0
datas = {}
for h3 in h3s:
    index = index + 2
    name = h3.select_one("a")["name"]

    table = tables[index]
    rows = table.findAll("tr")[1:]

    tableData = []
    for row in rows:
        data = row.findAll(['td', 'th'])
        cycles = convertCycles(data[3].get_text().strip())
        tableData.append({
            "name": name,
            "mode": convertMode(data[0].get_text()),
            "opcode": convertOpcode(data[1].get_text()),
            "bytes": int(convertBytes(data[2].get_text())),
            "cycles": int(cycles[0]),
            "pageCycle": int(cycles[1]),
        })

    datas[name] = tableData

with open('instructions.json', 'w') as file:
    json.dump(datas, file, ensure_ascii=False, indent=2)
