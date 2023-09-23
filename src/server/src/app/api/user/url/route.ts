import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {
    console.log("req url init")
    const d = await request.json()
    const { license } = d

    const { data, error } = await supabase
    .from('users')
    .select('auth_url')
    .eq('license', license)

    if (error) {
        console.log("couldn't reach url from db")
        return NextResponse.json({ message: 'error deleting user' }, { status: 500 })
    }

    const auth_url = data[0].auth_url
    
    return NextResponse.json({ auth_url:  auth_url}, { status: 200 })
}