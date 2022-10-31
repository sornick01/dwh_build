package main;


import java.sql.*;
import java.util.ArrayList;

import static org.jooq.impl.DSL.*;
import org.jooq.*;
import org.jooq.impl.*;

public class Main {

    public void createDstTables() {
        System.out.println("abc");
    }

    public static String query = "select * from \"payments\"";

    public static void main(String[] args) {
        try {
            Class.forName("org.postgresql.Driver");
            Connection srcCon = DriverManager.getConnection("jdbc:postgresql://localhost:5432/old", "postgres",
                    "Be8Bkf3z");
            Connection dstCon = DriverManager.getConnection("jdbc:postgresql://localhost:5432/new", "postgres",
                    "Be8Bkf3z");

//            Statement srcStatement = srcCon.createStatement();
//            Statement dstStatement = dstCon.createStatement();
//            while (rs.next()) {
//                System.out.println(rs.getString("title"));
//            }
            DSLContext dsl = DSL.using(srcCon, SQLDialect.POSTGRES);
            ArrayList<String> pk = new ArrayList<>();
            pk.add("a");
            pk.add("b");
            Attribute a = new Attribute("a", SQLDataType.VARCHAR, "newTable");
            Attribute b = new Attribute("b", SQLDataType.INTEGER, "newTable");
            Attribute c = new Attribute("c", SQLDataType.NUMERIC(14, 2), "newTable");
            ArrayList<Attribute> attrs = new ArrayList<Attribute>();
            attrs.add(a);
            attrs.add(b);
            attrs.add(c);
            Table table = new Table("newTable", attrs, pk, null);
            table.createTable(dsl);
            srcCon.close();
            dstCon.close();
        } catch (Exception e) {
            System.out.println(e);
        }
    }
}
