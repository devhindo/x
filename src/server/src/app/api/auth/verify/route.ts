import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {
    const data = await request.json()

    const license = data.license

    if(!license) {
        return NextResponse.json({ message: 'no license attached with the request'}, { status: 500 })
    }
    
    return await verify_license(license)
}

async function verify_license(l: string) {
    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('license', l)

    if (error || !data) {
        return NextResponse.json({ message: 'user is not registered yet'}, { status: 500 })
    }
    
    const access_token = data[0].access_token
    
    if ( !access_token ) {
        return NextResponse.json({ message: 'you haven\'t authorized X CLI yet'}, { status: 501 })
    }

    return NextResponse.json({ message: "verified. post a tweet using 'x -t \"first tweet from terminal!\""}, { status: 200 })
}

export async function GET(request: Request) {
    return new Response("Hello world")
}