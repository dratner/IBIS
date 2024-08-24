<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Phone Number Submission Form</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body {
            background-color: tan;
        }
        .bg-white a {
            text-decoration: underline;
        }
        .title {
            font-size: 1.25rem;
            font-weight: bold;
        }
    </style>
</head>
<body class="bg-tan min-h-screen flex flex-col">
    <div class="flex-grow flex items-center justify-center">
        <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
                    <img src="/static/IBIS.png" alt="Logo" class="mx-auto mb-6">

             <p class="text-blue-700 mb-6 text-center title">
                About IBIS
            </p>
            <p class="text-blue-700 text-sm mb-6 text-left">
                The Injured Bird Information System is a service operated by 
                <a href="https://snapdragonpartners.com">Snapdragon Partners</a> in partnership with <a href="https://www.birdmonitors.net/">Chicago Bird Collision Monitors</a>.
            </p>
            <p class="text-blue-700 text-sm mb-6 text-left">
                The purpose of the service is to efficiently route SMS (text message) alerts to volunteers in the network so that they can respond to requests to help injured birds.
            </p>
            <p class="text-blue-700 text-sm mb-6 text-left">
                In order to become a volunteer, please contact Chicago Bird Collision Monitors. If you are already a volunteer, you can <a HREF="/optin">opt in</a> for text alerts.
            </p>
            <p class="text-blue-700 text-sm mb-6 text-left">
                Once enrolled, you can reply STOP at any time to stop recieving text messages.
            </p>
        </div>
    </div>
    <footer class="bg-blue-600 text-white py-4">
        <div class="container mx-auto text-center">
            <a href="#" class="mx-2 hover:underline">About</a>
            <a href="#" class="mx-2 hover:underline">Terms of Use</a>
            <a href="#" class="mx-2 hover:underline">Privacy Policy</a>
        </div>
    </footer>
</body>
</html>