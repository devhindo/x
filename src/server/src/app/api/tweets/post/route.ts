import { Client } from "twitter-api-sdk";
import { AuthClient } from "twitter-api-sdk/dist/types";
import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

type Tweet = {
    state?: string,
    tweet?: string
}

export async function POST(request: Request) {

    const client = new Client("#");

    const post = await client.tweets.createTweet(
        {
            text: "foo",
        }
    )

    //return NextResponse.json({ state, tweet })
}

export async function GET(request: Request) {
    return new Response("hello get tweet")
}