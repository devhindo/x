import { Client } from "twitter-api-sdk";
import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'
import { generate_CONFIDENTIAL_CLIENT_AUTH_HEADER } from "../../auth/route";

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

type Tweet = {
    license?: string,
    tweet?: string
}

async function verify_license(license: string) {
    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('license', license)

    if (error || !data) {
        return false
    }
    return true
}

export async function POST(request: Request) {

    const data = await request.json()
    const { tweet, license}: Tweet = data

    const license_exist = await verify_license(license as string)
    
    if (!license_exist) {
        return NextResponse.json({ message: 'license not found' }, { status: 500 })
    }

    const access_token = await get_access_token(license as string)
    
    const client = new Client(access_token as string);
    
       const post = await client.tweets.createTweet(
               {
                   text: tweet as string,
               }
    )

    return NextResponse.json({ message: 'tweet posted!' }, { status: 200 })
}

async function get_access_token(license: string) {

    const [refresh_token_exist, refresh_token] = await get_refresh_token(license)

    if (!refresh_token_exist) {
        return
    }

    const [new_access_token, new_refresh_token] = await POST_new_access_token(license, refresh_token)
    
    await save_new_access_token(license, new_access_token, new_refresh_token)

    return new_access_token
}
    
async function save_new_access_token(license: string, access_token: string, refresh_token: string) {
    const { error } = await supabase
        .from('users')
        .update({ access_token: access_token, refresh_token: refresh_token })
        .eq('license', license)
    
    if (error) {
        return
    }      
}

async function update_refresh_token(license: string, refresh_token: string) {
    const { error } = await supabase
        .from('users')
        .update({ refresh_token: refresh_token })
        .eq('license', license)

    if (error) {
        return
    }
}

async function get_refresh_token(license: string): Promise<[boolean,string]> {
    const {data, error} = await supabase
        .from('users')
        .select()
        .eq('license', license)
        .single()
    if (error || !data) {
        return [false,""]
    } else {
        const { refresh_token } = data
        return [true,refresh_token as string]   
    }
}

async function POST_new_access_token(license: string, refresh_old_token: string): Promise<[string, string]> {

    const data = new URLSearchParams()
    data.append('grant_type', 'refresh_token')
    data.append('refresh_token', refresh_old_token)
    data.append('client_id', process.env.CLIENT_ID as string)

    const response = await fetch('https://api.twitter.com/2/oauth2/token', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': 'Basic ' + generate_CONFIDENTIAL_CLIENT_AUTH_HEADER(),
        },
        body: data
    })

    const json = await response.json()

    const { access_token, refresh_token } = json

    if(!access_token || !refresh_token) {
        console.log("new access token not found")
    }

    return [access_token, refresh_token]
}

export async function GET(request: Request) {
    return new Response("hello get tweet")
}