import os
import discord
from dotenv import load_dotenv

intents = discord.Intents.default()
client = discord.Client(intents = intents)

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')



@client.event
async def on_ready():
    print(f'{client.user} has connected to Discord!')

client.run(TOKEN)