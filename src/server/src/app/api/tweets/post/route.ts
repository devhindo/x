import { Client } from "twitter-api-sdk";
import { AuthClient } from "twitter-api-sdk/dist/types";
import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'
import { access } from "fs";

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

type Tweet = {
    license?: string,
    tweet?: string
}

export async function POST(request: Request) {

    const data = await request.json()
    const { tweet, license}: Tweet = data

    const access_token = await get_access_token(license as string)


    const client = new Client(access_token);

    const post = await client.tweets.createTweet(
        {
            text: tweet as string,
        }
    )

    return NextResponse.json({ message: 'tweet posted!' }, { status: 200 })}



async function get_access_token(license: string) {
    console.log("licessssssssssssssssssssssssssssssss" + license)
    const { data, error } = await supabase
        .from('users')
        .select()
        .eq('license', license)
        .single()
    if (error) {
        return NextResponse.json({ error, message: 'err' }, { status: 401 })
    }
    const { access_token } = data

    return access_token
}

export async function GET(request: Request) {
    return new Response("hello get tweet")
}