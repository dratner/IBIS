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
    </style>
</head>
<body class="bg-tan min-h-screen flex items-center justify-center">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <img src="/api/placeholder/200/100" alt="Logo" class="mx-auto mb-6">
        <form class="space-y-4">
            <div>
                <label for="phone" class="block text-sm font-medium text-blue-700">Phone Number</label>
                <input type="tel" id="phone" name="phone" required 
                       class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50">
            </div>
            <div class="flex items-start">
                <div class="flex items-center h-5">
                    <input id="agree" name="agree" type="checkbox" required
                           class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded">
                </div>
                <div class="ml-3 text-sm">
                    <label for="agree" class="font-medium text-blue-700">I agree to be contacted by IBIS</label>
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
</body>
</html>