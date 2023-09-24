import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {

    const d = await request.json()
    const license = d.license

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

    const auth_url = data[0].auth_url
    
    return NextResponse.json({ auth_url:  auth_url}, { status: 200 })
}