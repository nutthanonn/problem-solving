import cv2

img = cv2.imread('cat.jpeg', 0)
imgre = cv2.resize(img, (400, 200))


cv2.imshow('Cat', imgre)
cv2.waitKey(delay=5000)
cv2.destroyAllWindows()
