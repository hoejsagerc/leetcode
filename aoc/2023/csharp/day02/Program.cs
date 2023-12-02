using System.Text.RegularExpressions;

const string file = @"C:\Users\chris\dev\repos\leetcode\aoc\2023\csharp\day02\input.txt";
using var reader = new StreamReader(file);
var result1 = 0;
var result2 = 0;

void Main(StreamReader reader)
{
    while (reader.ReadLine()! is { } line)
    {
        var minimumCubes = new Dictionary<string, int>()
        {
            { "red", 0 }, { "green", 0 }, { "blue", 0 }
        };

        var gamePossible = true;

        var sets = line.Split(":")[1].Split(";");
        foreach (var set in sets)
        {
            var shownCubes = new Dictionary<string, int>()
            {
                { "red", 0 }, { "green", 0 }, { "blue", 0 }
            };
            foreach (var x in set.Split(","))
            {
                CheckValues(minimumCubes, shownCubes, x, "red");
                CheckValues(minimumCubes, shownCubes, x, "green");
                CheckValues(minimumCubes, shownCubes, x, "blue");
            }
            if (shownCubes["red"] > 12 || shownCubes["green"] > 13 || shownCubes["blue"] > 14)
            {
                gamePossible = false;
            }
        }
        result1 = CheckResult(result1, line, gamePossible);

        result2 += minimumCubes["red"] * minimumCubes["green"] * minimumCubes["blue"];
    }

    Console.WriteLine(result1);
    Console.WriteLine(result2);
}


static void CheckValues(Dictionary<string, int> minimumCubes,
    Dictionary<string, int> shownCubes, string x, string color)
{
    var pattern = @"^(\d+)";
    var newValue = int.Parse(Regex.Match(x.Trim(), pattern).Value);
    if (x.Contains(color))
    {
        shownCubes[color] = shownCubes[color] + newValue;

        if (minimumCubes[color] < newValue)
        {
            minimumCubes[color] = newValue;
        }
    }
}


static int CheckResult(int result1, string line, bool gamePossible)
{
    if (gamePossible)
    {
        var gameIndex = int.Parse(line.Split(":")[0].Split(" ")[1]);
        result1 += gameIndex;
    }

    return result1;
}


Main(reader);