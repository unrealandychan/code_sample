import httpx
import asyncio
import time

BASE_URL = "https://pokeapi.co/api/v2/pokemon/{}"


# For the Async function
async def get_request_async(number):
    async with httpx.AsyncClient() as client:
        return await client.get(BASE_URL.format(number))


pokemon_numbers = [i for i in range(1, 152, 1)]


async def fetch_pokemon():
    response = await asyncio.gather(*map(get_request_async, pokemon_numbers))
    data = [r.json()["species"]["name"] for r in response]
    return data


now = time.time()
async_pokemon = asyncio.run(fetch_pokemon())
print("Time consumed Async: {} seconds".format(time.time() - now))


# For loop functions
def get_requests_sync(number):
    response = httpx.get(BASE_URL.format(number))
    return response


now = time.time()
pokemons = []
for i in range(1, 152, 1):
    response = get_requests_sync(i)
    pokemons.append(response.json()["species"]["name"])
print("Time consumed Sync: {} seconds".format(time.time() - now))

assert async_pokemon == pokemons

# Output
"""
Time consumed Async: 1.383061170578003 seconds
Time consumed sync: 19.607528924942017 seconds

Obviously the async request is 19 times faster.
"""
