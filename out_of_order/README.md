# Out of Order Deserialization

This example demonstrates how to deserialize values from a byte slice in 
reverse order.

This is achieved by using the `Skip` method to find the offsets of values 
without fully decoding them.