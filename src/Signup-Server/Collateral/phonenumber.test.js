

//  Test phone --



TestCase("4155525686", "(415) 552-5686")
TestCase("(415) 552-5686", "(415) 552-5686")
TestCase("(415) 552-5686", "(415) 552-5686")
TestCase("(415) 552-568666666", "(415) 552-5686")
TestCase("#$*^&415dgdg", "(415) ");
TestCase("#$*^&4f1g5gdgdg", "(415) ");
TestCase("1",               "(1");
TestCase("12",              "(12");
TestCase("123",             "(123) ");
TestCase("1234",            "(123) 4");
TestCase("12345",           "(123) 45");
TestCase("123456",          "(123) 456-");
TestCase("1234567",         "(123) 456-7");
TestCase("12345678",        "(123) 456-78");
TestCase("123456789",       "(123) 456-789");
TestCase("1234567890",

TestCase("1 415 552 5686", "(415) 552-5686")

