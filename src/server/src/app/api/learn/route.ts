export async function GET(request: Request) {
    return new Response("learn page | GET");
}

// handle POST requests
export async function POST(request: Request) {
    const { username, secret } = await request.json();
    console.log(username, secret);
    return new Response("learn page | POST");
}