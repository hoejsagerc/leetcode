using System.Diagnostics.SymbolStore;
using System.Text.RegularExpressions;

const string file = @"C:\Users\chris\dev\repos\leetcode\aoc\2023\csharp\day01\input.txt";
using var reader = new StreamReader(file);

var numbers = new List<string>()
{
    "one","two","three","four","five","six","seven","eight","nine"
};



var result1 = 0;
var result2 = 0;
while (reader.ReadLine()! is { } line)
{
    var listOfDigits = new List<int>();
    var map = new Dictionary<int, int>();
    var charArray = line.ToCharArray();

    for(int i = 0; i < charArray.Count(); i++)
    {
        if (Char.IsDigit(charArray[i]))
        {
            map.Add(i, int.Parse(charArray[i].ToString()));
        }
    }
    result1 += int.Parse($"{map.OrderBy(x => x.Key).First().Value}{map.OrderBy(x => x.Key).Last().Value}");

    foreach(var num in numbers)
    {
        Regex regex = new Regex(num);
        MatchCollection matches = regex.Matches(line);

        foreach(Match match in matches)
        {
            map.Add(match.Index +1, MatchNumber(match.Value));
        }
    }

    result2 += int.Parse($"{map.OrderBy(x => x.Key).First().Value}{map.OrderBy(x => x.Key).Last().Value}");
}
Console.WriteLine(result1);
Console.WriteLine(result2);


int MatchNumber(string number) =>
    number switch
    {
        "one" => 1, "two" => 2, "three" => 3, "four" => 4, "five" => 5,
        "six" => 6, "seven" => 7, "eight" => 8, "nine" => 9,
        _ => throw new ArgumentException("Failed finding pattern matching",
            nameof(number))
    };