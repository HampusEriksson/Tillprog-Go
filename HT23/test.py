import turtle

# Create a turtle screen
screen = turtle.Screen()

# Create a turtle object
pen = turtle.Turtle()

# Set the background color of the screen
screen.bgcolor("white")

# Set the pen color using hexadecimal RGB values (Red)
pen.color(1, 0, 1)  # Red color in hexadecimal

pen.begin_fill()  # Begin filling the shape

# Draw a square
for _ in range(4):
    pen.forward(100)  # Move the turtle forward by 100 units
    pen.right(90)  # Turn the turtle right by 90 degrees

pen.end_fill()  # End filling the shape

# Close the turtle graphics window when clicked
screen.exitonclick()
