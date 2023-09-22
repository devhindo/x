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

    const { code_verifier, code_challenge, state_db } = user

    // request access token from twitter api
    //const access_token = await ax(code,code_verifier)
    const access_token = await req_access_token(code, code_verifier, state)
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
}


// Create a single supabase client for interacting with your database


const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)



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


/*
curl --location --request POST 'https://api.twitter.com/2/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'code=VGNibzFWSWREZm01bjN1N3dicWlNUG1oa2xRRVNNdmVHelJGY2hPWGxNd2dxOjE2MjIxNjA4MjU4MjU6MToxOmFjOjE' \
--data-urlencode 'grant_type=authorization_code' \
--data-urlencode 'client_id=rG9n6402A3dbUJKzXTNX4oWHJ' \
--data-urlencode 'redirect_uri=https://www.example.com' \
--data-urlencode 'code_verifier=challenge'
*/

// todo: https://twittercommunity.com/t/trying-to-get-oauth-2-0-token-receiving-missing-valid-authorization-header-error/163633/3 || 'Authorization: Basic CONFIDENTIAL_CLIENT_AUTH_HEADER'
// either here or here

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

// todo send the code verfier properly

async function req_access_token(code: string, verfier: string, state: string) {
    console.log("xxxxxxxxxxxxxxx" + verfier)
    let url = 'https://api.twitter.com/2/oauth2/token'
    url += '?grant_type=authorization_code'
    url += '&client_id=' + process.env.CLIENT_ID
    url += '&client_secret=' + process.env.CLIENT_SECRET
    url += '&redirect_uri=' + 'http://localhost:3000/api/auth'
    url += '&code=' + code
    url += '&code_verifier=' + verfier

    const {searchParams} = new URL(url)
    const ver = searchParams.get("code_verifier")
    console.log("2222222222222222" + ver)
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': 'Basic ' + generate_CONFIDENTIAL_CLIENT_AUTH_HEADER(),
        },
        //body: JSON.stringify({
        //    grant_type: 'authorization_code',
        //    client_id: process.env.CLIENT_ID as string,
        //    client_secret: process.env.CLIENT_SECRET as string,
        //    redirect_uri: 'http://localhost:3000/api/auth/validate/success',
        //    code: code,
        //    code_verifier: verfier,
        //})
    })
    const json = await response.json()
    console.log(json)
    return json
}
