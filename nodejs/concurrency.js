const axios = require("axios")


const BASE_URL =  "https://pokeapi.co/api/v2/pokemon/"



const get_pokemon_name_async = async (number)=>{
     const reponse = await axios.get(BASE_URL+i)
     return await reponse.data.species.name
}



const collect_first_gen_pokemon = async ()=>{
    let pokemon_async = []
    console.time()
    for (i=1; i <=151;i++){
        pokemon_async.push(get_pokemon_name_async(i))
    }
    const pokemons = await Promise.all(pokemon_async)
    console.log(pokemons)
    console.timeEnd()
}

collect_first_gen_pokemon()
