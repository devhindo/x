import { NextResponse } from "next/server"
import { createClient } from '@supabase/supabase-js'

export async function GET(request: Request) {
    console.log("start of requesting access token")
    const { searchParams } = new URL(request.url)
    
    const state = searchParams.get("state")
    const code = searchParams.get("code")

    if (!state || !code) {
        return new NextResponse("Missing state or code | Authentication failed", { status: 400 })
    }

    // todo: request access token from twitter api

    // todo: check if state is valid
    if (!await check_if_state_valis(state)) {
        return new NextResponse("Invalid state | Authentication failed", { status: 400 })
    }

    const user = await get_user_data(state)

    const { code_verifier } = user

    // request access token from twitter api
    //const access_token = await get_access_token(code,code_verifier)
    const access_token = await req_access_token(code,code_verifier)
    console.log("access_token:" + access_token)

    // todo: save all this in supabase

    // todo: response with access token to the cli upon post request

    // return access token
    return NextResponse.json({ access_token })
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

type User = {
    state?: string
    code_verifier?: string
    code_challenge?: string
}

async function check_if_state_valis(state: string) {
    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('state', state)
    .maybeSingle()

    if (error) {
        console.log(error)
    }
    if(data) {
        return true
    }
    return false
}

async function get_user_data(state: string) {
    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('state', state)
    .maybeSingle()

    if (error) {
        console.log(error)
    }
    if(data) {
        return data
    }
    return null
}

import axios from 'axios'

async function get_access_token(code: string, code_verfier: string) {
    const url = 'https://api.twitter.com/oauth2/token'
    const data = new URLSearchParams()
    data.append('code', code)
    data.append('grant_type', 'authorization_code')
    data.append('client_id', process.env.TWITTER_CLIENT_ID as string)
    data.append('redirect_uri', process.env.TWITTER_REDIRECT_URI as string)
    data.append('code_verifier', code_verfier)

    axios.post(url, data, {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
    .then(response => {
        console.log(response.data)
    }
    )
    .catch(error => {
        console.log(error)
    }
    )
}

// insert user
/*
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
*/
// insertUser(user)





/*
curl --location --request POST 'https://api.twitter.com/2/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'code=VGNibzFWSWREZm01bjN1N3dicWlNUG1oa2xRRVNNdmVHelJGY2hPWGxNd2dxOjE2MjIxNjA4MjU4MjU6MToxOmFjOjE' \
--data-urlencode 'grant_type=authorization_code' \
--data-urlencode 'client_id=rG9n6402A3dbUJKzXTNX4oWHJ' \
--data-urlencode 'redirect_uri=https://www.example.com' \
--data-urlencode 'code_verifier=challenge'
*/
/*
async function get_access_token(code: string) {
    const data = {
        code: code,
        grant_type: 'authorization_code',
        client_id: process.env.TWITTER_CLIENT_ID,
        redirect_uri: process.env.TWITTER_REDIRECT_URI,
        code_verifier: process.env.TWITTER_CODE_VERIFIER
    }

    const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: JSON.stringify(data)
    }

    fetch('https://api.twitter.com/2/oauth2/token', options)
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.log(error))


    return NextResponse.json({ data })
}
*/ 

async function req_access_token(code: string, code_verfier: string) {
    const url = 'https://api.twitter.com/2/oauth2/token'
    const headers = {
        'Content-Type': 'application/x-www-form-urlencoded',
    }
    const data = {
        grant_type: 'authorization_code',
        code: code,
        client_id: process.env.TWITTER_CLIENT_ID,
        redirect_uri: process.env.TWITTER_REDIRECT_URI,
        code_verifier: code_verfier
    }

    const options = {
        method: 'POST',
        headers: headers,
        body: JSON.stringify(data)
    }

    const response = await fetch(url, options)
    const json = await response.json()
    console.log(json)
    return json
}