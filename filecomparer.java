import java.io.File;
import java.util.Scanner;

public class FileComparer {
	public static void main(String[] args) {
		if(args.length != 2)
			System.our.println("Requires two files");
		File f1 = new File(args[0]);
		File f2 = new File(args[1]);
		Scanner sc1 = new Scanner(f1);
		Scanner sc2 = new Scanner(f2);
		int index = 0;
		while(sc1.hasNextLine() && sc2.hasNextLine()){
			String l1 = sc1.nextLine();
			String l2 = sc2.nextLine();
			if(!l1.equals(l2)){
				System.out.println(l1 + " != " + l2);
				SYstem.out.println("	=> Line #" + index);
				index++;
			}
		}
		if(sc1.hasNextLine()){
			System.out.println("FIle #1 still has data")
		}
		if(sc2.hasNextLine()){
			System.out.println("FIle #2 still has data")
		}
	}
}