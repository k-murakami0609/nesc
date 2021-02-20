opcodes := map[string]Opcode{}
{% for row in rows %}
opcodes[{{row["opcode"]}}] = Opcode{Name: {{row["name"]}}, Mode: Mode{{ row["mode"] }}, Code: {{ row["opcode"] }}, Cycle: {{ row["cycles"] }}, Size: {{ row["bytes"] }}, PageCycle: {{ row["pageCycle"] }}}{% endfor %}


