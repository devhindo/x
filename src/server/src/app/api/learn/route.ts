export async function GET(request: Request) {
    // wait 10 seconds
    // console log comming requests
    //console.log(request);
    console.log("___________________________________");
    console.log(request.body);
    console.log("___________________________________");
    console.log(request.headers);
    return new Response("learn page");
}