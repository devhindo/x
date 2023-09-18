import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'


export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    const state = searchParams.get('state')
    const code = searchParams.get('code')

    if (!state || !code) {
        return new Response("Missing state or code", { status: 400 })
    }
    
    const res = await fetch(`https://api.twitter.com/2/oauth2/token`, {
        headers: {
          'Content-Type': 'application/json',
          'code': code,
          'grant_type': 'authorization_code',
          'client_id' : 'id',
          'redirect_uri' : 'redirect',
          'code_verifier' : 'verifier'
        },
      })

    
    
}


// Create a single supabase client for interacting with your database


const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

// insert user

interface User {
    username: string,
    access_token: string
}

let user = {} as User
user.username = 'test'
user.access_token = 'test'

const insertUser = async (user: any) => {
    const { data, error } = await supabase.from('users').insert([
        { username: user.username, access_token: user.access_token},
    ])
    if (error) {
        console.log(error)
    }
    console.log(data)
}

console.log("ahmeddddddddd")
// insertUser(user)