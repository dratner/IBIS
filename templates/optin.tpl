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
            <p class="text-blue-700 mb-6 text-center title">
                Opt In
            </p>
            <p class="text-blue-700 text-sm mb-6 text-center">
                Thank you for volunteering for the IBIS Project to help injured birds. When you submit this form you will be entered into our volunteer database and you'll receive SMS notifications of birds you can assist.
            </p>
            <form class="space-y-4">
                <div>
                    <label for="phone" class="block text-sm font-medium text-blue-700">Phone Number</label>
                    <input type="tel" id="phone" name="phone" required 
                           class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50">
                </div>
                <div class="flex items-start">
                    <div class="flex items-center h-5">
                        <input id="agree" name="agree" type="checkbox" required
                               class="focus:ring-blue-500 h-4 w-4 text-blue-600 border border-gray-300 rounded">
                    </div>
                    <div class="ml-3 text-sm">
                        <label for="agree" class="font-medium text-blue-700">I agree to be contacted by IBIS (operated by Snapdragon Partners.)</label>
                    </div>
                </div>
                <div>
                    <button type="submit" 
                            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                        Submit
                    </button>
                </div>
            </form>
        </div>
    </div>
    <footer class="bg-blue-600 text-white py-4">
        <div class="container mx-auto text-center">
            <a href="/" class="mx-2 hover:underline">About</a>
            <a href="#" class="mx-2 hover:underline">Terms of Use</a>
            <a href="#" class="mx-2 hover:underline">Privacy Policy</a>
        </div>
    </footer>
</body>
</html>