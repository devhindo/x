import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {
    const data = await request.json()
    const license = data.license
    console.log(license)
    return await verify_license(license)
}

async function verify_license(license: string) {
    const { data, error } = await supabase
    .from('users')
    .select()
    .eq('license', license)

    if (error || !data) {
        console.log("no userrrrrrrrrr")
        return NextResponse.json({ error, message: 'err' }, { status: 500 })
    }

    // check if data is empty
    console.log("successsssssssssssss")
    return NextResponse.json({ message: 'verified'}, { status: 200 })
}

export async function GET(request: Request) {
    return new Response("Hello world")
}