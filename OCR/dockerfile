FROM python:3.10



RUN apt-get update && apt-get install -y && apt install -y tesseract-ocr tesseract-ocr-kor libgl1-mesa-glx
RUN pip3 install flask 
RUN apt install tesseract-ocr
RUN apt-get install tesseract-ocr-kor
RUN pip3 install pytesseract
RUN pip3 install imutils
RUN pip3 install numpy
RUN pip3 install opencv-python
RUN pip3 install scipy
RUN apt-get install libgl1-mesa-glx
RUN pip3 install matplotlib
RUN pip3 install requests
WORKDIR /app

COPY . /app

CMD ["python3", "-m", "flask", "run", "--host=0.0.0.0"]