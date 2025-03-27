const std = @import("std");
const print = std.debug.print;

fn middle(a: []const i32) i32 {
    var mid: i32 = 0;
    if (a.len % 2 == 0) {
        mid = a.len / 2;
    } else {
        mid = (a.len / 2) + 1;
    }
    return mid;
}

pub fn main() !void {
    const a = [_]i32{ 4, 6, 7, 9, 15, 20, 24, 26, 27, 29, 31 };

    var low: usize = 0;
    var high: usize = a.len - 1;

    var mid_index: usize = @divFloor(low + high, 2);
    var mid_value: i32 = a[mid_index];

    const search_num: i32 = 27;

    while (search_num != mid_value) {
        if (search_num > mid_value) {
            low = mid_index + 1;
        } else {
            high = mid_index - 1;
        }

        if (low > high) {
            print("Number not found.\n", .{});
            return;
        }

        mid_index = @divFloor(low + high, 2);
        mid_value = a[mid_index];
    }
    print("Index of {} is: {} \n", .{ search_num, mid_index });
}
