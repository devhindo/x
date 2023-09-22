import { NextResponse } from "next/server"
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

export async function GET(request: Request) {
    console.log("start of requesting access token")
    const { searchParams } = new URL(request.url)

    const state = searchParams.get("state")
    const code = searchParams.get("code")

    if (!state || !code) {
        return new NextResponse("Missing state or code | Authentication failed", { status: 400 })
    }

    if (!await check_if_state_valis(state)) {
        return new NextResponse("Invalid state | Authentication failed", { status: 400 })
    }

    const user = await get_user_data(state)

    const { code_verifier, code_challenge, state_db } = user

    //const access_token = await req_access_token(code, code_verifier, state)
    const [access_token, refresh_token, expires_in] = await req_access_token(code, code_verifier, state)

    if (await insert_access_token(state, access_token, refresh_token, expires_in)) {
        console.log("access token inserted")
    } else {
        console.log("access token not inserted")
    }
    //console.log(code_verifier)

    // todo: save all this in supabase

    // todo: response with access token to the cli upon post request

    // return access token
    return NextResponse.json({ access_token })
}


type User = {
    state?: string
    code_verifier?: string
    code_challenge?: string
    license?: string
}


// Create a single supabase client for interacting with your database






async function check_if_state_valis(state: string) {
    const { data, error } = await supabase
        .from('users')
        .select()
        .eq('state', state)
        .maybeSingle()

    if (error) {
        console.log(error)
    }
    if (data) {
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
    if (data) {
        return data
    }
    return null
}



function generate_CONFIDENTIAL_CLIENT_AUTH_HEADER() {
    const CLIENT_ID = process.env.CLIENT_ID
    const CLIENT_SECRET = process.env.CLIENT_SECRET
    return Buffer.from(CLIENT_ID + ":" + CLIENT_SECRET).toString('base64')
}

/*
POST https://authorization-server.com/token

grant_type=authorization_code
&client_id=i1zzqeXYK_4ZeAEU5_u6b8Qf
&client_secret=UaOzV9MrAijJMDmhucCBRo2XrRlPIfqBg1b0HCXePsDyqm88
&redirect_uri=https://www.oauth.com/playground/authorization-code-with-pkce.html
&code=HV24toek3g0qCYg3zZZcC4lSjE4IlI3nAZBRAK3thKyGeVQT
&code_verifier=mdUIOYqdl9r5EnV0hAoB6zznhaxqGTY-rXu-jQRwpDL5BE86
*/

// todo make the req more secure by performing a porpper post request
async function req_access_token(code: string, verfier: string, state: string): Promise<[string, string, number]> {
    let url = 'https://api.twitter.com/2/oauth2/token'
    url += '?grant_type=authorization_code'
    url += '&client_id=' + process.env.CLIENT_ID
    url += '&client_secret=' + process.env.CLIENT_SECRET
    url += '&redirect_uri=' + 'http://localhost:3000/api/auth'
    url += '&code=' + code
    url += '&code_verifier=' + verfier

    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': 'Basic ' + generate_CONFIDENTIAL_CLIENT_AUTH_HEADER(),
        },
    })

    const json = await response.json()
    console.log(json)
    // extract access_token and refresh_token
    const { access_token, refresh_token, expires_in } = json

    return [access_token, refresh_token, expires_in]
}

async function insert_access_token(state: string, access_token : string, refresh_token: string, expires_in: number) {
    const { data, error } = await supabase
    .from('users')
    .upsert({state: state ,access_token: access_token, refresh_token: refresh_token, expires_in: expires_in})
    .select()
    if (error) {
        console.log("errrrrrrrrrrrrr")
        console.log(error)

        return false
    } 
    if (data) {
        console.log("sucessssssssss")
        console.log(data)
        return true
    }
}

async function add_data_to_supabase(key: string, value: string, state: string) {
    const { error } = await supabase
    .from('users')
    .update({key: value})
    .eq('state', 'state')
    if (error) {
        console.log("couldn't insert " + key + error)
    }
}