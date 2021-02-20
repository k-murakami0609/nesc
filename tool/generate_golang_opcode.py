import json
import jinja2
from itertools import chain


with open("instructions.json") as f:
    df = json.load(f)

    templateLoader = jinja2.FileSystemLoader(searchpath="./")
    templateEnv = jinja2.Environment(loader=templateLoader)
    template = templateEnv.get_template("template.tpl")

    rows = list(chain.from_iterable(df.values()))
    dispText = template.render({"rows": rows})
    print(dispText)
