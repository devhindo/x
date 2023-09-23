import { Client } from "twitter-api-sdk";
import { AuthClient } from "twitter-api-sdk/dist/types";
import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'
import { access } from "fs";
import { verify_license } from "../../auth/verify/route";
import { generate_CONFIDENTIAL_CLIENT_AUTH_HEADER } from "../../auth/route";

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

type Tweet = {
    license?: string,
    tweet?: string
}

export async function POST(request: Request) {

    const data = await request.json()
    const { tweet, license}: Tweet = data

    const license_exist = await verify_license(license as string)
    
    if (!license_exist) {
        return NextResponse.json({ message: 'license not found' }, { status: 500 })
    }

    const access_token = await get_access_token(license as string)

    console.log("uppppppppppppperaccess token: " + access_token)
    
    const client = new Client(access_token as string);
    
       const post = await client.tweets.createTweet(
               {
                   text: tweet as string,
               }
    )

    return NextResponse.json({ message: 'tweet posted!' }, { status: 200 })
}

        

async function get_access_token(license: string) {

    const [refresh_token_exist, refresh_token] = await get_referesh_token(license)

    if (!refresh_token_exist) {
        console.log("refresh token not found")
        return
    }

    const [new_access_token, new_referesh_token] = await POST_new_access_token(license, refresh_token)
    
    console.log("55new access token: " + new_access_token)
    console.log("55new referesh token: " + new_referesh_token)

    await save_new_access_token(license, new_access_token, new_referesh_token)

    return new_access_token
}
    
async function save_new_access_token(license: string, access_token: string, referesh_token: string) {
    var { error } = await supabase
        .from('users')
        .update({ access_token: access_token })
        .eq('license', license)
    
    if (error) {
        console.log("couldn't update new access token" + error)
        return
    }      
    
    var {error} = await supabase
        .from('users')
        .update({ refresh_token: referesh_token })
        .eq('license', license)
    
    if (error) {
        console.log("couldn't update new referesh token" + error)
        return
    }

    console.log("new access token saved")
}
    // I have referesh token

    //const { data, error } = await supabase
    //    .from('users')
    //    .select()
    //    .eq('license', license)
    //    .single()
    //if (error) {
    //    return NextResponse.json({ error, message: 'err' }, { status: 401 })
    //}
    //const { access_token } = data
//
    //return access_token


async function get_referesh_token(license: string): Promise<[boolean,string]> {
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

async function POST_new_access_token(license: string, referesh_old_token: string): Promise<[string, string]> {

    const data = new URLSearchParams()
    data.append('grant_type', 'refresh_token')
    data.append('refresh_token', referesh_old_token)
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

    console.log(json)



    const { access_token, referesh_token } = json

    console.log("newwwwwwwwtoken: " + access_token)
    console.log("newwwwwwwwreferesh: " + referesh_token)
    if(!access_token || !referesh_token) {
        console.log("new access token not found")
    }

    return [access_token, referesh_token]
}

export async function GET(request: Request) {
    return new Response("hello get tweet")
}