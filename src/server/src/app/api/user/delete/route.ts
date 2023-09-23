import { NextResponse } from "next/server";
import { createClient } from '@supabase/supabase-js'

const supabase = createClient(process.env.SUPABASE_URL as string, process.env.SUPABASE_SECRET as string)

export async function POST(request: Request) {
    
    const data = await request.json()
    const { license } = data

    const { error } = await supabase
    .from('users')
    .delete()
    .eq('license', license)

    if (error) {
        return NextResponse.json({ message: 'error deleting user' }, { status: 500 })
    }

    return NextResponse.json({ message: 'user deleted!' }, { status: 200 })
}