import requests
import os

url = "https://discord.com/api/v8/applications/862056701553147965/commands"
quickTestUrl = "https://discord.com/api/v8/applications/862056701553147965/guilds/862322123539087380/commands"

jsons = [
    {
        "name": "blot",
        "description": "Managing battles!",
        "options": [
            {
                "type": "1",
                "name": "bloop",
                "description": "Place your team",
                "options": [
                    {
                        "type": 4,
                        "name": "x",
                        "description": "The X coordinate your team will be placed",
                        "required": True,
                    },
                    {
                        "type": 4,
                        "name": "y",
                        "description": "The Y coordinate your team will be placed",
                        "required": True,
                    },
                ]
            },
            {
                "type": "1",
                "name": "ploint",
                "description": "Point your team",
                "options": [
                    {
                        "type": 4,
                        "name": "angle",
                        "description": "The angle in degrees your team will attempt to march",
                        "required": True,
                    },
                ]
            },
            {
                "type": "1",
                "name": "clolour",
                "description": "Choose your team's colour",
                "options": [
                    {
                        "type": 3,
                        "name": "colour",
                        "description": "The angle in degrees your team will attempt to march",
                        "required": True,
                    },
                ]
            },


        ]
    }
]
# For authorization, you can use either your bot token
headers = {
    "Authorization": "Bot " + os.environ['BOT_TOKEN']
}

# We use PUT here because it overwrites
r = requests.put(quickTestUrl, headers=headers, json=jsons)
print(r.text)
