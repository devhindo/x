import { Client } from "twitter-api-sdk";
import { AuthClient } from "twitter-api-sdk/dist/types";
import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'
import { access } from "fs";

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

type Tweet = {
    state?: string,
    tweet?: string
}

export async function POST(request: Request) {

    const data = await request.json()
    const tweet: Tweet = data


    const access_token = await get_access_token(tweet.state as string)

    const client = new Client(access_token);

    const post = await client.tweets.createTweet(
        {
            text: tweet.tweet as string,
        }
    )

    return new NextResponse("tweet posted!")
}



async function get_access_token(state: string) {
    const { data, error } = await supabase
        .from('users')
        .select()
        .eq('state', state)
        .single()
    if (error) {
        console.log(error)
    }
    const { access_token } = data

    return access_token
}

export async function GET(request: Request) {
    return new Response("hello get tweet")
}