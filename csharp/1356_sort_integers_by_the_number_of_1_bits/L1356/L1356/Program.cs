
using System.Collections;


var myArray = new int[] {1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1};

SortByBits(myArray);

int[] SortByBits(int[] arr)
{
    var map = new List<Dictionary<int, int>>();

    foreach (var i in arr)
    {
        Console.WriteLine("");
        Console.WriteLine($"Checking bits in number: {i}");
        BitArray barray = new BitArray(new int[] { i });
        var count = 0;
        foreach (var bit in barray)
        {
            if (bit.Equals(true))
            {
                Console.Write("1 ");
                count++;
            }
            else
            {
                Console.Write("0 ");
            }
        }

        map.Add(new Dictionary<int, int>(){{ i, count }});
    }

    // var sortedList = map.Select(dict => dict.OrderBy(kvp => kvp.Value).ToDictionary(kvp => kvp.Key, kvp => kvp.Value));
    var sortedList = map.OrderBy(x => x.Values.Max()).ThenBy(x => x.Keys.FirstOrDefault()).ToList();
    var resultArray = new List<int>();
    foreach (var x in sortedList)
    {
        resultArray.Add(x.Keys.FirstOrDefault());
    }

    return resultArray.ToArray();
}
