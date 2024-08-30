export async function analyzeImage(imageUrl) {
    const prompt = import.meta.env.VITE_PROMPT;

    console.log('Prompt:', prompt);

    const response = await fetch('https://api.openai.com/v1/chat/completions', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${import.meta.env.VITE_OPENAI_API_KEY}`,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            model: 'gpt-4o-mini',
            messages: [
                {
                    role: 'system',
                    content: prompt,
                },
                {
                    role: 'user',
                    content: [{
                        type: 'image_url',
                        image_url: {
                            url: imageUrl,
                        },
                    }]
                }
            ],
            max_tokens: 800,
        }),
    });

    const data = await response.json();
    console.log(data);
    return data.choices[0].message.content.trim();
}
