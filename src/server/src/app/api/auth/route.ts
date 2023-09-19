import { NextResponse } from "next/server"
import { createClient } from '@supabase/supabase-js'

export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    
    const state = searchParams.get("state")
    const code = searchParams.get("code")

    if (!state || !code) {
        return new NextResponse("Missing state or code | Authentication failed", { status: 400 })
    }

    // todo: request access token from twitter api

    // todo: save all this in supabase

    // todo: response with access token to the cli upon post request

    return NextResponse.json({ state, code })
}

//type User = {
//    username?: string
//    secret?:string
//}
//
//export async function POST(request: Request) {
//    const data: User = await request.json()
//    console.log(data)
//
//    const { username, secret } = data
//
//    return NextResponse.json({ username, secret })
//}


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