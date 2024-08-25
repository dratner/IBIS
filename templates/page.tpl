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
        {{ .Content }}
         </div>
    </div>
    <footer class="bg-blue-600 text-white py-4">
        <div class="container mx-auto text-center">
        <div class="mb-2">
            <a href="/" class="mx-2 hover:underline">About</a>
            <a href="/terms" class="mx-2 hover:underline">Terms of Use</a>
            <a href="/privacy" class="mx-2 hover:underline">Privacy Policy</a>
            </div>
            <div class="text-xs">
            Copyright &copy; 2024 by Snapdragon Partners, LLC. All rights reserved.
        </div>
        </div>
    </footer>
     <script>
        const phoneInput = document.getElementById('phone');
        const agreeCheckbox = document.getElementById('agree');
        const submitButton = document.getElementById('submitButton');

        function validateForm() {
            const phoneValid = phoneInput.value.trim() !== '';
            const agreeChecked = agreeCheckbox.checked;

            if (phoneValid && agreeChecked) {
                submitButton.disabled = false;
                submitButton.classList.remove('bg-blue-300', 'cursor-not-allowed');
                submitButton.classList.add('bg-blue-600', 'hover:bg-blue-700');
            } else {
                submitButton.disabled = true;
                submitButton.classList.add('bg-blue-300', 'cursor-not-allowed');
                submitButton.classList.remove('bg-blue-600', 'hover:bg-blue-700');
            }
        }

        phoneInput.addEventListener('input', validateForm);
        agreeCheckbox.addEventListener('change', validateForm);
    </script>
</body>
</html>
