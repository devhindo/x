import { NextResponse } from 'next/server'
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string , process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {
    const data = await request.json()
    const license = data.license
    console.log(license)

    const license_exist = await verify_license(license)

    if (!license_exist) {
        return NextResponse.json({ message: 'license not found'}, { status: 500 })
    }

    return NextResponse.json({ message: 'verified'}, { status: 200 })

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

export async function GET(request: Request) {
    return new Response("Hello world")
}