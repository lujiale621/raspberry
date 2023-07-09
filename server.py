from fastapi import FastAPI, HTTPException
import requests

app = FastAPI()

@app.post("/text-to-speech")
async def text_to_speech(data: Dict[str, str]):
    try:
        	text = data["text"]   
    except KeyError as e:
        	return {"err_code": 1002, "err_desc": e, "result": None}
    url = "http://localhost:7860/voice/vits"
    data = {
        "text": text,
        "id": 0,
        "format": "wav",
        "lang": "auto",
        "length": 1.2,
        "noise": 0.667,
        "noisew": 0.8,
        "max": 50
    }

    try:
        response = requests.post(url, data=data)
        response.raise_for_status()
        
        with open("output.wav", "wb") as f:
            f.write(response.content)
        
        return {"message": "success"}
    
    except requests.exceptions.RequestException:
        raise HTTPException(status_code=500, detail="Request to voice service failed")

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
