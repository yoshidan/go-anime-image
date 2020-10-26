import cv2
import sys
import os.path

def detect(cascade_file = './lbpcascade_animeface.xml'):
    if not os.path.isfile(cascade_file):
        raise RuntimeError('%s: not found' % cascade_file)

    cascade = cv2.CascadeClassifier(cascade_file)

    indir = './download/'
    outdir = './face'
    for filename in os.listdir(indir):
        if filename == '.gitkeep':
            continue
        if os.path.exists('%s/0_%s' % (outdir, filename)):
            continue

        image = cv2.imread(indir + filename, cv2.IMREAD_COLOR)
        gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
        gray = cv2.equalizeHist(gray)

        faces = cascade.detectMultiScale(gray,scaleFactor = 1.1, minNeighbors = 5, minSize = (48, 48))
        for i in range(len(faces)):
            (x, y, w, h) = faces[i]
            new_image = image[y:y+h , x:x+w]
            path = '%s/%s_%s' % (outdir, str(i), filename)
            print('save as ' + path)
            cv2.imwrite(path, new_image)

detect()