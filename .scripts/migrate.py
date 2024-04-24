import csv
import json
import requests

with open("t_games.csv", encoding='utf-8') as r_file:
    file_reader = csv.reader(r_file, delimiter=",")
    for row in file_reader:
        body = json.dumps({
            "done": row[1] == 'true',
            "name": row[2],
            "image": row[5],
            "hltb_id": int(row[6]),
            "clear_path_image": True,
        })

        response = requests.post("http://localhost:8080/api/v1/game", data=body, headers={"Content-Type": "application/json"})
        print(response.status_code, "|", row[2])