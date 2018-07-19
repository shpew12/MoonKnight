from random import seed, random
from time import clock
from operator import itemgetter
from collections import namedtuple
from math import sqrt
from copy import deepcopy
import numpy as np
import bisect
import pdb

def sqd(p1, p2):
    return sum((c1 - c2) ** 2 for c1, c2 in zip(p1, p2))




def binary_search(a, x, lo=0, hi=None):  # can't use a to specify default for hi
    hi = hi if hi is not None else len(a)  # hi defaults to len(a)
    pos = bisect.bisect_left(a, x, lo, hi)  # find insertion position
    return (pos if pos != hi and a[pos] == x else -1)  # don't walk off the end



class KdNode(object):
    __slots__ = ("dom_elt", "split", "left", "right","inxbound", "bound")

    def __init__(self, dom_elt, split, left, right, inxbound, bound):
        self.dom_elt = dom_elt
        self.split = split
        self.left = left
        self.right = right
        self.bound = bound
        self.inxbound = inxbound
class KdTree(object):
    __slots__ = ("n","equalize","resolution","pts", "bounds")
    def __init__(self, pts, bounds, resolution, equalize =False):
        pdb.set_trace()
        self.equalize = equalize
        #add bounds based off of random_points
        #takes pts and puts them into d-dimensional list of sorted coordinates with associated points
        y = np.array([[(j,jnx) for jnx, j in enumerate(i)]  for i in np.array(pts).T ])
        self.pts = np.array([i[i[:,0].argsort()] for i in y])

# min  to  min+1/2(max-min)
# max - 1/2(max-min)  to  max

        def nk2(split,resolution,inxbound,bound):
            if resolution <=0 :
                return None

            lcbound = deepcopy(bound)
            rcbound = deepcopy(bound)
            lcinxbound=deepcopy(inxbound)
            rcinxbound=deepcopy(inxbound)


            middle = bound[split][0]+(bound[split][1]-bound[split][0])/2.0
            #change upper bound for left child and lower bound for right child
            lcbound[split][1] = middle
            rcbound[split][0] = middle


            inxmiddle = bisect.bisect([i[0] for i in self.pts[split]], middle , lo=inxbound[split][0] , hi=inxbound[split][1])
            #change the upper index for the left child and lower index for right child
            lcinxbound[split][1]=inxmiddle
            rcinxbound[split][0]=inxmiddle

            resolution =resolution -1


            s2 = (split + 1) % len(self.pts)  # cycle coordinates
            return KdNode(middle, split, inxbound[split], bound[split], nk2(s2, resolution,lcinxbound,lcbound ),
                                    nk2(s2, resolution,rcinxbound,rcbound ))

        inxbnd = [[0,len(pts)] for x in range(len(self.pts))]
        self.n = nk2(0,resolution,inxbnd,self.bounds)




T3 = namedtuple("T3", "nearest dist_sqd nodes_visited")


def find_nearest(k, t, p):
    def nn(kd, target, hr, max_dist_sqd):
        if kd is None:
            return T3([0.0] * k, float("inf"), 0)

        nodes_visited = 1
        s = kd.split
        pivot = kd.dom_elt
        left_hr = deepcopy(hr)
        right_hr = deepcopy(hr)
        left_hr.max[s] = pivot[s]
        right_hr.min[s] = pivot[s]

        if target[s] <= pivot[s]:
            nearer_kd, nearer_hr = kd.left, left_hr
            further_kd, further_hr = kd.right, right_hr
        else:
            nearer_kd, nearer_hr = kd.right, right_hr
            further_kd, further_hr = kd.left, left_hr

        n1 = nn(nearer_kd, target, nearer_hr, max_dist_sqd)
        nearest = n1.nearest
        dist_sqd = n1.dist_sqd
        nodes_visited += n1.nodes_visited

        if dist_sqd < max_dist_sqd:
            max_dist_sqd = dist_sqd
        d = (pivot[s] - target[s]) ** 2
        if d > max_dist_sqd:
            return T3(nearest, dist_sqd, nodes_visited)
        d = sqd(pivot, target)
        if d < dist_sqd:
            nearest = pivot
            dist_sqd = d
            max_dist_sqd = dist_sqd

        n2 = nn(further_kd, target, further_hr, max_dist_sqd)
        nodes_visited += n2.nodes_visited
        if n2.dist_sqd < dist_sqd:
            nearest = n2.nearest
            dist_sqd = n2.dist_sqd

        return T3(nearest, dist_sqd, nodes_visited)

    return nn(t.n, p, t.bounds, float("inf"))


def show_nearest(k, heading, kd, p):
    print(heading + ":")
    print("Point:           ", p)
    n = find_nearest(k, kd, p)
    print("Nearest neighbor:", n.nearest)
    print("Distance:        ", sqrt(n.dist_sqd))
    print("Nodes visited:   ", n.nodes_visited, "\n")


def random_point(k):
    return [random() for _ in range(k)]


def random_points(k, n):
    return [random_point(k) for _ in range(n)]

if __name__ == "__main__":
    seed(1)

    #P = lambda *coords: list(coords)


    pts = [[1,14],[2,17],[3,19],[4,15],[5,16],[6,21],
    [7,23],[8,18],[9,22],[10,24],[11,20],[12,26],[13,25]]
    bnds=np.array([[0.0, 14.0], [13.0, 27.0]])
    print(bnds)
    print(pts)

    kd1 = KdTree(pts,bnds,5)





    '''
    show_nearest(2, "Wikipedia example data", kd1, P(9, 2))

    N = 400000
    t0 = clock()
    kd2 = KdTree(random_points(3, N), Orthotope(P(0, 0, 0), P(1, 1, 1)))
    t1 = clock()
    text = lambda *parts: "".join(map(str, parts))
    show_nearest(2, text("k-d tree with ", N,
                         " random 3D points (generation time: ",
                         t1-t0, "s)"),
                 kd2, random_point(3))
    '''
